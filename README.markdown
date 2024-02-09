The following code helps me to understand how to use complex domain in PostgreSQL using [pgx v5](https://github.com/jackc/pgx).

In this case I used an array of another domain:

```SQL
CREATE DOMAIN public.http_client_error_code AS smallint
	CONSTRAINT http_client_error_code_check CHECK (VALUE = 0 OR VALUE >= 400 AND VALUE <= 499);
```

The above domain is a simple integer with a constraint.

```SQL
CREATE DOMAIN public.http_client_error_code_list AS http_client_error_code[];
```

The above domain is an array of the first domain (`http_client_error_code`).

When trying to scan or assign it using pgx an error will be raised:

```
panic: failed to encode args[0]: unable to encode main.HTTPClientErrorCodeList{400} into text format for unknown type (OID 32783): cannot find encode plan
```

The issue is that `pgx` does not understand how to convert the slice into `http_client_error_code_list`.

By telling `pgx` to look for that domain (and given OID), it know how to `scan` and `value` the Go's data type.

