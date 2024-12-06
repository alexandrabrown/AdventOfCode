from itertools import pairwise

def is_strictly_increasing(my_list):
    if all(a < b for a, b in pairwise(my_list)):
        return True
    else:
        return False

def is_strictly_decreasing(my_list):
    if all(a > b for a, b in pairwise(my_list)):
        return True
    else:
        return False
    
def check_differences(my_list):
    if all(abs(a-b) >=1 and abs(a-b) <=3 for a, b in pairwise(my_list)):
        return True
    else:
        return False
    
def is_safe(my_list):
    return (is_strictly_increasing(my_list) or is_strictly_decreasing(my_list)) and check_differences(my_list)

def main():
    numSafe = 0
    with open('input.txt', 'r') as file:
        for line in file:
            my_list = [float(x) for x in line.split()]
            if (is_safe(my_list)):
                numSafe += 1

    print(numSafe)


if __name__ == "__main__":
    main()