
# URL Shortener
Simple url shortener generator

## Run Locally

Clone the project

```bash
  git clone git@github.com:fwidyatama/shortener-pizza.git
```

Go to the project directory

```bash
  cd shortener
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  make running-local
```


## Endpoint

| Endpoint  	         | Method   	 | Body / Query (Example)  	                                                          |
|---------------------|------------|------------------------------------------------------------------------------------|
| /api/shorten	       | POST   	   | ``` {"destination":"https://www.google.com","expire_at":"2023-09-12 15:00:00"} ``` |
| /api/{short_url}  	 | GET   	    | -    	                                                                             |
| /api/list  	        | GET   	    | ?order=DESC <br/> ?order=ASC   	                                                   |

