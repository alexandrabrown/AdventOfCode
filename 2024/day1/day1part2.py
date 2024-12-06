def read_file_word_by_word(filename):
    with open(filename, 'r') as file:
        for line in file:
            for word in line.split():
                yield word

def main():
    list1 = list();
    list2 = list();
    for index, word in enumerate(read_file_word_by_word('input.txt')):
        if (index % 2):
            list2.append(int(word))
        else:
            list1.append(int(word))
    result = 0
    for num1 in list1:
        for num2 in list2:
            if (num1 == num2):
                result += num1
    print(result)

if __name__ == "__main__":
    main()