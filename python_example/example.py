"""
This app connects to the de-challenge API and retrieves the songs from 2021-01-01.
"""

import json
import os
import requests

DAILY_SONGS_ENDPOINT = "https://de-challenge.ltvco.com/v1/songs/daily"
RELEASED_AT = "2021-01-01"
OK = 200


def main():
    body = {"api_key": os.getenv("API_KEY"), "released_at": RELEASED_AT}
    response = requests.get(DAILY_SONGS_ENDPOINT, params=body)
    print(json.dumps((response.json()), indent=4, sort_keys=False, ensure_ascii=False))


if __name__ == '__main__':
    main()
