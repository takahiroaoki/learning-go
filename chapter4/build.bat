echo Building synonyms...

cd synonyms
go build -o synonyms.exe

echo Building sprinkle...
cd ../sprinkle
go build -o sprinkle.exe

echo Building coolify...
cd ../coolify
go build -o coolify.exe

echo Building domainify...
cd ../domainify
go build -o domainify.exe

echo Building available...
cd ../available
go build -o available

echo Build done!