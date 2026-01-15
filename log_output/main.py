#!/usr/bin/python3

from datetime import datetime
from uuid import uuid4
from time import sleep

def main():
    random_string = uuid4()

    while True:
        time = datetime.now()
        print(f"{time}: {random_string}")
        sleep(5)

if __name__ == "__main__":
    main()
