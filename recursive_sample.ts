import * as fs from 'fs';

function globPaths(targetPath: string, maxDepth: number, suffix: string, depth = 0): string[] {
    const matchedPaths: string[] = [];

    if (depth === maxDepth) {
        console.log(`Reached maximum depth of ${maxDepth} at path ${targetPath}`);
    } else if (fs.statSync(targetPath).isDirectory()) {
        depth += 1;

        for (const listEntry of fs.readdirSync(targetPath)) {
            const subPath = `${targetPath}/${listEntry}`;
            const subPathMatchedPaths = globPaths(subPath, maxDepth, suffix, depth);
            matchedPaths.push(...subPathMatchedPaths);
        }
    } else if (fs.statSync(targetPath).isFile()) {
        if (targetPath.endsWith(suffix)) {
            matchedPaths.push(targetPath);
        }
    }

    return matchedPaths;
}

function main() {
    const result = globPaths(process.cwd(), 2, ".json");
    for (const item of result) {
        console.log(item.replace(/\.json$/, ".json5"));
    }
}

main();
