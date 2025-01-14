# How to run

- Enter inside root folder of the project;
- Type de command ```go mod tidy``` to make sure the dependencies of project is ok;
- Type de command ```make``` to generate Swagger docs and start the server on
port ```3333``` or another port of your preference;
- Type the address in your browser ```http://localhost:3333/docs/index.html``` to
see the project Swagger docs.

If you want to test the API endpoins you should be install the VSCode extension [clicking here](https://marketplace.visualstudio.com/items?itemName=humao.rest-client), or any API client of your choice.

Enjoy :-)!

## Project dependencies

- Golang 1.23.4 or higher;
- make (to run the makefile tasks);
