import json
from dataclasses import dataclass
from typing import List

from .xor import *


@dataclass
class Event:
    lookup_id: str
    stamp_ns: int
    node: Key  # node performing the lookup
    target: Key  # target of the lookup
    request: StateUpdateEvent
    response: StateUpdateEvent

    def heard(self):
        req = self.request.heard if self.request else []
        resp = self.response.heard if self.response else []
        return req + resp

    def waiting(self):
        req = self.request.waiting if self.request else []
        resp = self.response.waiting if self.response else []
        return req + resp

    def queried(self):
        req = self.request.queried if self.request else []
        resp = self.response.queried if self.response else []
        return req + resp

    def unreachable(self):
        req = self.request.unreachable if self.request else []
        resp = self.response.unreachable if self.response else []
        return req + resp


@dataclass
class StateUpdateEvent:
    cause: Key  # peer that caused this event
    source: Key
    heard: List[Key]
    waiting: List[Key]
    queried: List[Key]
    unreachable: List[Key]


def parse_state_update_from_json(data):
    if not data:
        return None
    else:
        return StateUpdateEvent(
            cause=key_from_base64(data["Cause"]["Kad"]),
            source=key_from_base64(data["Source"]["Kad"]),
            heard=[key_from_base64(id["Kad"]) for id in data.get("Heard", [])],
            waiting=[key_from_base64(id["Kad"]) for id in data.get("Waiting", [])],
            queried=[key_from_base64(id["Kad"]) for id in data.get("Queried", [])],
            unreachable=[key_from_base64(id["Kad"]) for id in data.get("Unreachable", [])],
        )


def parse_event_from_json(data):
    return Event(
        lookup_id=data["ID"],
        stamp_ns=data["XXX"],  # XXX
        node=key_from_base64(data["Node"]["Kad"]),
        target=key_from_base64(data["Key"]["Kad"]),
        request=parse_state_update_from_json(data.get("Request")),
        response=parse_state_update_from_json(data.get("Response")),
    )


def load_file(filename: str):
    events = []
    with open(filename) as f:
        for line in f:
            data = json.loads(line)
            events.append(parse_event_from_json(data))
    return events


def filter_lookup(events: List[Event], lookup_id: str):
    return [ev for ev in events if ev.lookup_id == lookup_id]
