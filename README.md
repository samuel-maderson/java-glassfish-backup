# TO-DO

Create a program that reads a json file, inside this json file we'll have the absolute path for
backup both directory application on domains/ of glassfish directory and MySQL path of dumped file.

1. First dump the MySQL and safe the dump file anywhere else
2. Compress the application directory in the same place where our dumped file was stored.
3. Upload both files in a AWS S3 Bucket.

Remember, all parameters must be set on a json file, this file is the where our program will look for to know what he have to do.

OBS: Credentials will be passed as program argument.