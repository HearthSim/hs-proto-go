if [ ! -w "${GOPATH}" ]; then
	echo "\$GOPATH must be set"
	exit 1
fi
if [ ! "${GOPATH}/src/github.com/HearthSim/hs-proto-go" -ef "$(pwd)" ]; then
	echo "hs-proto-go should be in \$GOPATH"
	exit 1
fi
if [ ! -x "$(which protoc)" ]; then
	echo "protoc not found"
	exit 1
fi
if [ ! -x "${GOPATH}/bin/protoc-gen-go" ]; then
	echo "Use `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}` to acquire protoc-gen-go"
	exit 1
fi
if [ ! -f "${HSPATH}/Assembly-CSharp.dll" ]; then
	echo "\$HSPATH must be a path containing Assembly-CSharp.dll"
	exit 1
fi
if [ ! -f "${HSPATH}/Assembly-CSharp-firstpass.dll" ]; then
	echo "\$HSPATH must be a path containing Assembly-CSharp-firstpass.dll"
	exit 1
fi
if [ ! -r "${PROTOX}" ]; then
	echo "\$PROTOX must be set to the binary of <https://github.com/HearthSim/proto-extractor>"
	exit 1
fi

PROTO_OUT="${GOPATH}/src/github.com/HearthSim/hs-proto-go"
echo "Extracting .proto files to ${PROTO_OUT}"
"${PROTOX}" -g "${PROTO_OUT}" \
	"${HSPATH}"/Assembly-CSharp{,-firstpass}.dll
cd "${GOPATH}/src"
for i in github.com/HearthSim/hs-proto-go/{bnet,pegasus}/*/*.proto; do
	echo "Building ${i}"
	PATH="${GOPATH}/bin:${PATH}" protoc --go_out=. "${i}"
done
