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
    with open("input.txt") as f:
        for line in f:
            code_quest(line)

if __name__ == "__main__":
    main()