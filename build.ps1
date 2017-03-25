# Usage: PowerShell.exe -ExecutionPolicy Unrestricted .\build.ps1 "D:\Git\proto-extractor\bin\Debug\proto-out"
# Requirements:
#   The protobuffer compiler tool (https://github.com/google/protobuf/releases)
#   and add it to your path, '$GOPATH/bin' could be used!
#   PROTO and PROTOC-GEN-GO GO packages:
#       go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

param(
    # First positional parameter - The root directory of all .proto files.
    # Point this path to the output directory of the Proto-Extrator project!
    [string]$PROTO_IN,
    # Second positional parameter - GOPATH, if not set as environment variable
    [string]$GO_PATH
)

# Must be forward slashes so the compiler does not complain!
# This must also END WITH A SLASH.
Set-Variable -Name Import_Prefix -Value "github.com/HearthSim/hs-proto-go/"

# Check provided proto path.
if( (!$PROTO_IN) -or !(Test-Path $PROTO_IN) ) {
    Write-Host "ERROR   First parameter must be a path to the protobuffer files!"
    exit 1
}

# Check provided GOPATH, or use the environment variable.
if(!$GO_PATH) {
    $Go_Path_Env = Get-ChildItem Env:GOPATH

    if(!$Go_Path_Env) {
        Write-Host "ERROR   The environment variable GOPATH must be set or passed as second argument!"
        exit 1
    } else {
        $GO_PATH = $Go_Path_Env.Value
    }
}

# Update environment variable PATH to include the protocompiler extensions
$env:Path = "$env:Path;$GO_PATH\bin"

# Generate complete output path for the compiled proto files.
#  === $GOPATH/github.com/HearthSim/hs-proto-go/ === 
$Compiled_Proto_Path = Join-Path $GO_PATH -ChildPath "src" | Join-Path -ChildPath $Import_Prefix

Write-Host "INFO    Compiled protobuffer files are written to: "$Compiled_Proto_Path
Write-Host

# Create directory if it did not exist.
# Out-Null is used to pipe output of the command to a null stream.
New-Item -ItemType Directory -Force -Path $Compiled_Proto_Path | Out-Null
Set-Location $Compiled_Proto_Path

# Compile all proto files per package.
# So we take each subdirectory of the $PROTO_IN path and compile all
# .proto files found inside that directory at the same time.
$Packages = Get-ChildItem -Path $PROTO_IN -Recurse -Directory
$Packages | ForEach-Object {
    $Proto_Files = Get-ChildItem $_.FullName -Filter *.proto -File
    $PackageName = $_.Name.ToLower()

    if($Proto_Files.Length -gt 0) {
        $File_List = ""

        $Proto_Files | ForEach-Object {
            $File_List += "`"" + $_.FullName + "`"" 
        }

        Write-Host "INFO    Building package: "$_.FullName
        protoc --proto_path="$PROTO_IN" --go_out="import_prefix=${Import_Prefix}:." $File_List
        Write-Host
    }    
}

Write-Host "INFO    Initiating fixes.."

Write-Host " - Fixing greedy imports.."
# Fix for greedy prefixing of import statements by the GO compiler
Get-ChildItem -Path $Compiled_Proto_Path -Recurse -Filter *.go -File |
ForEach-Object {
    $File_Path = $_.FullName
    $Contents = (Get-Content $File_Path)

    # Do simple string replacement
    $Contents.Replace('"github.com/HearthSim/hs-proto-go/github.com/golang/protobuf/proto"', '"github.com/golang/protobuf/proto"') |
    Set-Content $File_Path
}

Write-Host "INFO    Process finished!"