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
    used: List[Key]  # used is a dict of all keys of peers that were attempted during the lookup
    queries: List[QueryModel]
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

    def find_closest_success_query(self):
        closest = None
        for q in self.queries:
            if q.outcome == QUERY_SUCCESS:
                if not closest or q.peer.to_float() < closest.peer.to_float():
                    closest = q
        return closest

    def find_source_query(self, peer):
        for q in self.queries:
            if peer in q.response.heard:
                return q
        return None

    def find_path(self):
        closest = self.find_closest_success_query()
        if not closest:
            return []
        else:
            path = []
            next = closest
            while next:
                path.append(next)
                next = self.find_source_query(next)
            return path


@dataclass
class QueryModel:
    peer: Key  # peer being queried
    request: Event  # query request event
    response: Event  # query response event
    outcome: str  # success, unreachable, unfinished


QUERY_SUCCESS = "success"
QUERY_UNREACHABLE = "unreachable"
QUERY_UNFINISHED = "unfinished"


def request_events(events):
    return [e for e in events if e.request]


def response_events(events):
    return [e for e in events if e.response]


def request_matches_response(req, resp):
    if len(req.request.waiting) == 1 and resp.response.cause == req.request.waiting[0]:
        return resp.response.cause
    else:
        return None


def queries_from_events(events):
    queries = []
    for req in request_events(events):
        for resp in response_events(events):
            peer = request_matches_response(req, resp)
            break
        if peer:
            queries.append(
                QueryModel(
                    peer=peer,
                    request=req,
                    response=resp,
                    outcome=QUERY_SUCCESS if len(resp.response.unreachable) == 0 else QUERY_UNREACHABLE,
                )
            )
        else:
            queries.append(
                QueryModel(
                    peer=peer,
                    request=req,
                    response=None,
                    outcome=QUERY_UNFINISHED,
                )
            )
    return queries


def lookup_from_events(events):
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
        queries=queries_from_events(events),
        events=events,
    )
