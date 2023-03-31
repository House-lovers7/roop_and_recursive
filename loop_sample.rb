require 'find'

# ターゲットパスを調整する
def adjust_path(target_path)
  if target_path.end_with?('/')
    target_path
  else
    target_path + '/'
  end
end

# 指定されたパスから、指定された最大深度以下の、指定された拡張子に一致するファイルパスを取得する
def glob_paths(target_path, suffix, max_depth)
  matched_paths = [] # 一致したファイルパスを格納するための配列
  target_path = adjust_path(target_path)
  Find.find(target_path) do |current_path|
    if File.directory?(current_path) # パスがディレクトリの場合
      if File.basename(current_path) == '.' || File.basename(current_path).start_with?('.')
        Find.prune # カレントディレクトリまたは隠しファイルの場合はスキップする
      else
        depth = current_path.count('/') - target_path.count('/')
        if depth >= max_depth # 指定された深度を超えたら処理を終了する
          Find.prune
        end
      end
    elsif File.file?(current_path) && current_path.end_with?(suffix) # パスがファイルで、指定された拡張子に一致する場合
      matched_paths << current_path # 一致したファイルパスを配列に追加する
    end
  end
  matched_paths # 一致したファイルパスの配列を返す
end

result = glob_paths('.', '.json', 3)
puts "JSONファイルの配列: #{result}"
