# GraphQL Example Using GQL Gen

The purpose of this example is to provide details as to how one would go about creating a Restful API with the Buffalo Web Framework.

## Getting Started

## Software requirements

- [Go 1.10.3 or higher](https://golang.org/dl)

## Communication

- If you **need help**, use [Stack Overflow](http://stackoverflow.com/questions/tagged/graphql). (Tag 'graphql')
- If you'd like to **ask a general question**, use [Stack Overflow](http://stackoverflow.com/questions/tagged/graphql).
- If you **found a bug**, open an issue.
- If you **have a feature request**, open an issue.
- If you **want to contribute**, submit a pull request.

## Quick Installation

1.  clone this repository

    ```text
    $ cd $GOPATH/src/github.com
    $ git clone git@github.com:conradwt/gqlgen_example.git
    ```

2.  change directory location

    ```text
    $ cp /path/to/gqlgen_example
    ```

3.  start the server

    ```text
    $ go run server/server.go
    ```

4.  navigate to the site in the browser

    ```text
    $ open http://localhost:8080
    ```

5.  test some queries

    ```text
    mutation createTodo {
      createTodo(input:{text:"todo", userId:"1"}) {
        user {
          id
        }
        text
        done
      }
    }

    query findTodos {
      todos {
        text
        done
        user {
          name
        }
      }
    }
    ```

## GQL Gen References

- Official website: https://gqlgen.com
- Guides: https://gqlgen.com
- Docs: https://gqlgen.com
- Source: https://github.com/99designs/gqlgen

## Support

Bug reports and feature requests can be filed with the rest for the GraphQL Example Using GQL Gen project here:

- [File Bug Reports and Features](https://github.com/conradwt/gqlgen_example/issues)

## License

GraphQL Example Using GQL Gen is released under the [MIT license](https://mit-license.org).

## Copyright

copyright:: (c) Copyright 2018 Conrad Taylor. All Rights Reserved.
