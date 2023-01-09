from utils import download_input

if __name__ == "__main__":
    data = download_input(6)

    def get_marker(x):
        for i in range(len(x)):
            if len(set(x[i : i + 14])) == 14:
                return i + 14
