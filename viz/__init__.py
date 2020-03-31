from .events import *
from .xor import *
from .model import *
from .plot import *

if __name__ == "__main__":
    ev = filter_lookup(
        load_file("/Users/petar/data/ad2c7194d190/cypress-searcher/0/dht_lookups.out"),
        "ed102f55-ebe5-475c-a528-99c050ba259f")
    print(len(ev))
