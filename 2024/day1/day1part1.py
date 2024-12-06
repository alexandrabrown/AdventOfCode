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
    list1.sort()
    list2.sort()
    result = [abs(x - y) for x, y in zip(list1, list2)]
    print(sum(result))

if __name__ == "__main__":
    main()