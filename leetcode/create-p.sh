# ./create-p.sh p00053 MaximumSubarray  
mkdir $1 && cd $1 && echo "package $1\n\n// $2 ...\nfunc $2(){\n}" > "$1.go" && ginkgo generate && ginkgo bootstrap && touch "$2.md"
