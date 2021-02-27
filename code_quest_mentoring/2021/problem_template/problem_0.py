import os 

DIR_PATH = os.path.dirname(os.path.realpath(__file__))
INPUT_FILE = DIR_PATH + "/input.txt"

def code_quest(number_string):
    number = int(number_string)
    if number % 21 == 0:
        print("CODEQUEST")
    elif number % 7 == 0:
        print("QUEST")
    elif number % 3 == 0:
        print("CODE")
    else:
        print(number)

def main():
    with open(INPUT_FILE) as f:
        for line in f:
            code_quest(line.strip())

if __name__ == "__main__":
    main()