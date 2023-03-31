require 'pathname'

def glob_paths(target_path, max_depth, suffix, depth = 0)
  matched_paths = []

  if depth == max_depth
    puts "Reached maximum depth of #{max_depth} at path #{target_path}"
  else
    if File.directory?(target_path)
      depth += 1

      Dir.foreach(target_path) do |list_item|
        next if list_item == '.' || list_item == '..'

        sub_path = File.join(target_path, list_item)
        sub_path_matched_paths = glob_paths(sub_path, max_depth, suffix, depth)
        matched_paths.concat(sub_path_matched_paths)
      end
    elsif File.file?(target_path)
      if File.extname(target_path) == suffix
        matched_paths << target_path
      end
    end
  end

  matched_paths
end

def main
  result = glob_paths('.', 2, '.json')
  result.each do |item|
    puts Pathname.new(item).dirname.join(Pathname.new("#{Pathname.new(item).basename('.*')}.json5"))
  end
end

main
