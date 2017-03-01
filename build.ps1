param(
    [string]$PROTO_OUT,
    [string]$GO_PROTO_OUT
)

# Change directory to the GO_PROTO repository
Set-Location $GO_PROTO_OUT

# Compile all proto files per package.
# So we take each subdirectory of the $PROTO_OUT path and compile all
# .proto files found inside that directory at the same time.
$Packages = Get-ChildItem $PROTO_OUT -Recurse -Directory
$Packages | ForEach-Object {
    $Proto_Files = Get-ChildItem $_.FullName -Filter "*.proto"

    if($Proto_Files.Length -gt 0) {
        $File_List = ""

        $Proto_Files | ForEach-Object {
            $File_List += "`"" + $_.FullName + "`"" 
        }

        Write-Host "Building package: "$_.FullName
        protoc --proto_path=$PROTO_OUT --go_out=. $File_List
        Write-Host
    }    
}