#!/usr/bin/env python3

import sys
import os

if __name__ == "__main__":
    argv = sys.argv

    argv = argv[1:]

    if len(argv) == 0:
        print("ERROR: no test file was provided", file=sys.stderr)
        sys.exit(1)

    test_file_name = argv[0]
    if test_file_name.endswith(".piled"):
        test_file_name = test_file_name.removesuffix(".piled")
    if test_file_name.endswith(".expected"):
        test_file_name = test_file_name.removesuffix(".expected")

    tests_dir = "./tests/"
    test_filepath = os.path.join(tests_dir, test_file_name)

    with open(test_filepath + ".piled", "w"):
        pass
    
    with open(test_filepath + ".expected", "w"):
        pass
    
    print(f"Create {test_filepath + '.piled'} Successfully.")
    print(f"Create {test_filepath + '.expected'} Successfully.")

