The docs are written in [MkDocs](https://www.mkdocs.org/) using the
[Material for MkDocs](https://squidfunk.github.io/mkdocs-material/) theme.

MkDocs is a static site generator converting Markdown files to static HTML pages.
Edit the markdown files, and view the rendered site with MkDocs.

## Local development and testing

### Install MkDocs and dependencies

To ensure you are using a compatible version of  we recommend to use a Python 
[virtual environment](https://docs.python.org/3/library/venv.html) to install 
`mkdocs-material`:

```
python3 -m venv venv
source venv/bin/activate
pip install mkdocs-material
```

### Local preview

Test locally, including both the content and the navigation structure.
The docs are just MD files and in order to see the rendered preview 
locally (before PRing the repo), do:

```
 $ mkdocs serve
INFO    -  Building documentation...
INFO    -  Cleaning site directory
INFO    -  Documentation built in 0.35 seconds
[I 210610 13:17:46 server:334] Serving on http://127.0.0.1:8000
INFO    -  Serving on http://127.0.0.1:8000
[I 210610 13:17:46 handlers:62] Start watching changes
INFO    -  Start watching changes
[I 210610 13:17:46 handlers:64] Start detecting changes
...
```

Some useful pointers:


* For issues with the local preview check the [Material for MkDocs changelog](https://squidfunk.github.io/mkdocs-material/upgrading/).
* The [MkDocs reference](https://www.mkdocs.org/user-guide/writing-your-docs/) includes information on the structure of the `docs/` folder, and MD authoring.
* Review the [Material for MkDocs reference](https://squidfunk.github.io/mkdocs-material/reference/formatting/) for information on the theme and formatting.

## Publishing

You don't need to do anything. Commits that change a file under the `docs/` 
directory trigger a [GitHub action](https://github.com/pollypkg/polly/blob/main/.github/workflows/update-docs.yaml)
to build the site and deploy it to GitHub pages.

