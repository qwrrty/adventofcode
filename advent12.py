#! /usr/bin/env python

import json

def findNumbers(thing, ignore_red=False):
    nums = []
    if isinstance(thing, basestring):
        pass
    elif isinstance(thing, int):
        nums.append(thing)
    elif isinstance(thing, list):
        for elem in thing:
            nums.extend(findNumbers(elem, ignore_red))
    elif isinstance(thing, dict):
        for k, v in thing.iteritems():
            if ignore_red and v == 'red':
                nums = []
                break
            nums.extend(findNumbers(v, ignore_red))
    else:
        raise TypeError(thing)
    return nums

def sumNumbers(struct, ignore_red=False):
    return sum(findNumbers(struct, ignore_red))
    
def main():
    with open("input12.json") as f:
        x = json.load(f)
    print sumNumbers(x, ignore_red=True)

if __name__ == '__main__':
    main()
