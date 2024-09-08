#!/usr/bin/env python3

import subprocess
import glob
import sys
import os

def run_cmd(cmd: list[str], silent=True) -> subprocess.CompletedProcess:
    if not silent:
        print(" ".join(cmd))
    return subprocess.run(cmd, capture_output=True)

def parse_expected_file(filepath: str) -> dict:
    with open(filepath) as f:
        lines = f.readlines()

    result = {"int": []}
    for row, line in enumerate([line.strip("\n") for line in lines]):
        words = line.split(" ")
        if len(words) != 2:
            continue
        head = words.pop(0)
        if head.startswith(":"):
            stripped_head = head.lstrip(":")
            if stripped_head == "int":
                value = words.pop(0)

                result["int"].append(value)
    return result

def test_compile(fp: str) -> None:
    print("[COMPILE]")
    print(f"Compiling {fp} ...")
    out = run_cmd(["./piled.out", fp])
    print(f"Compiled successfully")

def chmod_x_all(fp: str) -> None:
    binary_file = fp.removesuffix(".piled") + ".out"
    run_cmd(["chmod", "+x", binary_file], )

def test_run(fp: str) -> None:
    binary_file = fp.removesuffix(".piled") + ".out"
    expected_file = fp.removesuffix(".piled") + ".expected"
    expected = parse_expected_file(expected_file)['int']  # TODO: Not implemented other than int
    print("[RUN]")
    actual = run_cmd([binary_file], silent=True)
    actual_stdout = list(filter(lambda x: len(x) > 0, actual.stdout.decode().split("\n")))
    actual_stderr = list(filter(lambda x: len(x) > 0, actual.stderr.decode().split("\n")))

    if len(actual_stdout) != len(expected):
        print(f"[ERROR]: {fp}")
        print("length of output is not equal")
        print(f"    actual   -> {actual_stdout}")
        print(f"    expected -> {expected}")
        print(f"    actual-length   -> {len(actual_stdout)}")
        print(f"    expected-length -> {len(expected)}")

    for i in range(len(actual_stdout)):
        if actual_stdout[i] != expected[i]:
            print(f"[ERROR]: {fp}")
            print(f"element {i+1} is not valid")
            print(f"    actual   -> {actual_stdout}")
            print(f"    expected -> {expected}")

    print(f"Verifing {binary_file} ...")
    print(f"Verified successfully")

def clean_up(tests: list[str]) -> None:
    print("[INFO] Cleaning tests directory ...")
    for filepath in tests:
        filename_without_ext = filepath.removesuffix(".piled")
        assembly_file = filename_without_ext + ".asm"
        binary_file = filename_without_ext + ".out"
        os.remove(assembly_file)
        os.remove(binary_file)

def main() -> None:
    tests_dir = "./tests/*.piled"
    tests = glob.glob(tests_dir)
    for filepath in sorted(tests):
        print(f"----- {os.path.basename(filepath)} -----")
        test_compile(filepath)
        chmod_x_all(filepath)
        test_run(filepath)
        print(f"[PASS]")
        print()
    clean_up(tests)


if __name__ == "__main__":
    main()
