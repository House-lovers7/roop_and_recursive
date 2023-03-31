import * as fs from 'fs';
import * as path from 'path';

function adjustPath(targetPath: string): string {
    if (targetPath.endsWith('/')) {
        return targetPath;
    } else {
        return targetPath + '/';
    }
}

export function globPaths(targetPath: string, suffix: string, maxDepth: number): string[] {
    const matchedPaths: string[] = [];
    const targetPaths: [string, number][] = [[adjustPath(targetPath), 0]];

    while (targetPaths.length > 0) {
        const [currentPath, depth] = targetPaths.shift()!;
        if (depth >= maxDepth) {
            break;
        }
        const stat = fs.statSync(currentPath);
        if (stat.isDirectory()) {
            for (const item of fs.readdirSync(currentPath)) {
                const subPath = path.join(currentPath, item);
                targetPaths.push([subPath, depth + 1]);
            }
        } else if (stat.isFile() && currentPath.endsWith(suffix)) {
            matchedPaths.push(currentPath);
        }
    }
    return matchedPaths;
}

function main() {
    const result = globPaths(process.cwd(), '.json', 3);
    console.log('JSONファイルの配列:', result);
}

main();
