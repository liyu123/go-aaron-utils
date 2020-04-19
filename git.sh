cd /Users/aaron/study/go_src/go-aaron-utils/||exit

time=$(date '+%F %T')

msg=$1' '$time

branch=$(git symbolic-ref --short -q HEAD)

git status

git add .
git commit -m"${msg}"

git push origin "${branch}"