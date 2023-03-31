import os
from typing import List

def glob_paths(target_path: str, max_depth: int, *, suffix: str, depth: int = 0) -> List[str]:
    matched_paths: List[str] = []
    if depth == max_depth:
        pass
        print(f'引数からのディレクトリの深さ{depth}/{max_depth}階層、ファイルPATH:{target_path}')
    elif os.path.isdir(target_path):
        depth += 1
        for list_item in os.listdir(target_path):
            sub_path = os.path.join(target_path, list_item)
            subpath_matched_paths: List[str] = glob_paths(sub_path, max_depth, suffix=suffix, depth=depth)
            matched_paths.extend(subpath_matched_paths)
    elif os.path.isfile(target_path):
        if target_path.endswith(suffix):
            matched_paths.append(target_path)
    return matched_paths

def main():
    result = glob_paths(target_path=os.getcwd(), max_depth=2, suffix=".json")
    for item in result:
        print(os.path.splitext(item)[0] + ".json5")
        os.path
main()
