import os
from typing import List, Tuple

def adj_path(path: str):
    if path.endswith("/"):
        pass
    else:
        path = path + "/"
    return path

def glob_paths(target_path: str, suffix: str, *, max_depth: int) -> List[str]:

    matched_paths: List[str] = []
    loop_target_paths: List[Tuple[str, int]] = []
    target_path_adj: str = adj_path(target_path)
    depth: int = 0
    loop_target_paths.append((target_path_adj, depth))

    while loop_target_paths:
        loop_target_path, depth = loop_target_paths.pop(0)
        if depth >= max_depth:
            break
        print(f'loop_target_paths:{loop_target_paths} \n')
        if os.path.isdir(loop_target_path):
            for list_item in os.listdir(loop_target_path):
                sub_path = os.path.join(loop_target_path, list_item)
                loop_target_paths.append((sub_path, depth + 1))
        elif os.path.isfile(loop_target_path):
            if loop_target_path.endswith(suffix):
                matched_paths.append(loop_target_path)
    return matched_paths

def main():
    result = glob_paths(os.getcwd(), suffix=".json", max_depth=3)
    print(f'JSONファイルの配列:{result}')
main()
