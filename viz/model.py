from dataclasses import dataclass
from typing import List

from .events import *
from .xor import *


@dataclass
class LookupModel:
    id: str
    start_ns: int
    stop_ns: int
    target: Key
    # used is a dict of all keys of peers that were attempted during the lookup
    used: List[Key]
    events: List[events]

    def stamp_to_x(self, stamp_ns: int):
        """Return the x-axis value for a given nanosecond timestamp."""
        # return milliseconds since the first event in the lookup
        return (stamp_ns - self.start_ns) / 1000000.0

    def min_x(self):
        return 0.0

    def max_x(self):
        return (self.stop_ns - self.start_ns) / 1000000.0

    def key_to_y(self, key: Key):
        """Return the y-axis value for a given key."""
        return xor_key(self.target, key).to_float()

    def min_y(self):
        return 0.0

    def max_y(self):
        return 1.0

    def event_key_xy(self, event, key):
        return self.stamp_to_x(event.stamp_ns), self.key_to_y(key)


def model_from_events(events):
    if len(events) < 2:
        raise Exception("Not enough events to plot")

    did = {}
    used = []

    def push(x):
        if not x:
            return
        if not did.get(x):
            did[x] = True
            used.append(x)

    for e in events:
        if e.request:
            push(e.request.cause)
            push(e.request.source)
        if e.response:
            push(e.response.cause)
            push(e.response.source)
    return LookupModel(
        id=events[0].lookup_id,
        start_ns=events[0].stamp_ns,
        stop_ns=events[-1].stamp_ns,
        target=events[0].target,
        used=used,
        events=events,
    )
