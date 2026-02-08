MIGRATION_NAME=$1

LATEST_MIGRATION_VERSION=1

for file in migrations/*.sql; do
    version=$(basename "$file" | cut -d'_' -f1)
    version=$((10#$version))
    if [[ "$version" > "$LATEST_MIGRATION_VERSION" ]]; then
        LATEST_MIGRATION_VERSION="$version"
    fi
done

LATEST_MIGRATION_VERSION=$((10#$LATEST_MIGRATION_VERSION + 1))
LATEST_MIGRATION_VERSION=$(printf "%05d" $LATEST_MIGRATION_VERSION)

echo "MIGRATION_NAME: $MIGRATION_NAME"
echo "MIGRATION_VERSION: $LATEST_MIGRATION_VERSION"

FILE_NAME="migrations/${LATEST_MIGRATION_VERSION}_${MIGRATION_NAME}.sql"

touch "$FILE_NAME"
echo "Created migration file: $FILE_NAME"