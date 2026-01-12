#!/usr/bin/python3

from uuid import uuid4
from time import sleep

def main():
    random_string = uuid4()

    while True:
        print(random_string)
        sleep(5)

if __name__ == "__main__":
    main()
