from datetime import datetime
from uuid import uuid4
from time import sleep
import os

def write_log(content: str, path: str):
    with open(path, 'a') as file:
        timestamp = datetime.now()
        file.write(f"{timestamp}: {content}\n")

def main():
    env_path: str | None = os.environ.get('LOG_DIR')
    if env_path and env_path != "":
        log_dir = env_path
    else:
        log_dir = "./logs"
    os.makedirs(log_dir, exist_ok=True)
    file_path = f"{log_dir}/timestamps.log"
    print(f"Writing logs to: {file_path}")

    random_string = uuid4()

    while True:
        write_log(random_string, file_path)
        sleep(5)

if __name__ == "__main__":
    print("Starting timestamp logger...")
    main()
