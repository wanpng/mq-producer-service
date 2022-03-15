SRC_DIR=$(pwd)
DST_DIR=$(pwd)

echo "OS:" $OSTYPE "[compiling protocol buffers]"

protoc -I=$SRC_DIR --go_out=$GOPATH/src $SRC_DIR/protos/domain/*.proto
protoc -I=$SRC_DIR --go_out=$GOPATH/src $SRC_DIR/protos/service/*.proto
