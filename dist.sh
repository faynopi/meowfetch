#!/usr/bin/env sh

DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Check if script running from the repo directory
if [[ "$DIR" != "$PWD" ]]; then
    echo "[Error] Please execute the script within the repository directory."
    exit 1
fi

TARGETS="386 amd64 arm arm64"
VERSION="$(make opts| grep VERSION | cut -d':' -f 2-)"
BUILDDIR="$(make opts| grep BUILDDIR | cut -d':' -f 2-)"

[[ -d "$DIR/dist" ]] && rm -rf "$DIR/dist"
[[ -n $(find $DIR -type f -iname "Source code.*") ]] && rm "$DIR/Source code."*

cd ../
tar -cf "Source code.tar" "$(basename $DIR)"
mv "Source code.tar" $DIR
cd - > /dev/null
gzip "Source code.tar"

mkdir -p "$DIR/dist"
cp -f "$DIR/Source code.tar.gz" "$DIR/dist"

IFS=" "; for TARGET in $TARGETS; do
    echo "Building meowfetch $VERSION for $TARGET"
    GOARCH=$TARGET make clean build > /dev/null  || continue
    [[ -e "$BUILDDIR/meowfetch" ]] || continiue

    TD="$DIR/dist/meowfetch-${VERSION}-${TARGET}" # TD => TargetDir
    mkdir -p $TD

    SHA256SUM="$(sha256sum "${BUILDDIR}/meowfetch" | cut -d' ' -f1)"
    echo "$SHA256SUM meowfetch-$VERSION-$TARGET" >> "$DIR/dist/sha256sum.txt"

	cp -f ${BUILDDIR}/meowfetch "${TD}"
	chmod 755 ${TD}/meowfetch

    sed "s/{VERSION}/${VERSION}/g" < meowfetch.1 > "${TD}/meowfetch.1"
	chmod 644 ${TD}/meowfetch.1

    cd dist
    tar -cf "$(basename $TD).tar" "$(basename $TD)"
	gzip ${TD}.tar
    cd - > /dev/null

    rm -rf "$TD"

done
