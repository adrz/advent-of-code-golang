import re

import requests

url = "https://adventofcode.com/2022/day/8/input"
headers = {
    "cookie": "session=53616c7465645f5f8d2f634fc3937da2ce2e570f21e07e49e97dba1512d32047d7cecd367df2c5fc2214385db077466a8de77d9b738f927cb489fd4613af1de7"
}
data = requests.get(url, headers=headers).text.splitlines()
