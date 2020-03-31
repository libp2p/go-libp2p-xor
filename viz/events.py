import json
from dataclasses import dataclass
from typing import List

from .xor import *


@dataclass
class Event:
    lookup_id: str
    stamp_ns: int
    target: Key
    cause: Key
    source: Key
    heard: List[Key]
    waiting: List[Key]
    queried: List[Key]
    unreachable: List[Key]


def load_file(filename: str):
    events = []
    with open(filename) as f:
        for line in f:
            data = json.loads(line)
            events.append(
                Event(
                    lookup_id=data["eventID"],
                    stamp_ns=data["ts"],
                    target=key_from_base64(data.get("targetKad")),
                    cause=key_from_base64_optional(data.get("causeKad")),
                    source=key_from_base64_optional(data.get("sourceKad")),
                    heard=[key_from_base64(h) for h in data.get("heardKad", [])],
                    waiting=[key_from_base64(h) for h in data.get("waitingKad", [])],
                    queried=[key_from_base64(h) for h in data.get("queriedKad", [])],
                    unreachable=[key_from_base64(h) for h in data.get("unreachableKad", [])],
                )
            )
    return events


def filter_lookup(events: List[Event], lookup_id: str):
    return [ev for ev in events if ev.lookup_id == lookup_id]
