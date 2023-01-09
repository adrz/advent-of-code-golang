import os

import requests
from dotenv import load_dotenv

load_dotenv()
COOKIE = os.getenv("COOKIE")

HEADERS = {"cookie": f"session={COOKIE}"}


def download_input(day, year=2022):
    url = f"https://adventofcode.com/{year}/day/{day}/input"
    r = requests.get(url, headers=HEADERS)
    r.raise_for_status()
    return r.text
