# A Command line tool for deepl

**deepldoc**  is a tool that automatically translates your project's documentation into another language. Written in the Go language, it uses DeepL's API as the translation engine.

## How to use

### 1. Create a Free API Account

- After [creating an account](https://www.deepl.com/pro?cta=apiDocsHeader#developer), obtain an API key.
- Set the API key as an environment variable by executing the following command in the terminal:

```sh
export DEEPL_API_KEY=your_api_key_here
```

### 2. Using the `deepl` Command

- The `deepl` command is used as follows:

```sh
. /deepl text [target_language]
``` 

- **text** is the text to be translated.
- **target_language** is the language code of the target language.

### 3. Using the `deepldoc` Command

To use `deepldoc`, follow these instructions:

```sh
. /deepldoc source_directory [target_language] [file_extension]
```

- **source_directory** is the path to the directory containing the documents you want to translate.
- **target_language** is the language code of the target language (e.g. 'ja'). If omitted, 'ja' is used by default.
- **file_extension** is the extension of the file to be translated (e.g. 'md'). If omitted, 'md' is used by default.


### Notes

- **deepldoc** translates files with the relevant file extensions and copies files with non-relevant extensions as they are. This preserves the original directory structure.

- **deepldoc** preserves code blocks and inline code in your documents. Text wrapped in single (`) or triple backticks (```) will not be translated, ensuring accurate representation of your code samples.

## How to Download and Run Binaries

Follow these steps to download and run the `deepl` and `deepldoc` files:

### 1. Download

- Go to the [releases page](https://github.com/koriym/deepldoc/releases) and download `deepl` and `deepldoc`.

### 2. Make the Files Executable
- Open a terminal and navigate to your downloads folder:
```sh
cd ~/Downloads
```

- Make the files executable:

```sh
chmod +x deepl
chmod +x deepldoc
```

### 3. Run the files:

```sh
./deepl
./deepldoc
```

That's it!

## FAQ

### `error: unexpected status code: 403, body:` 

This indicates that the API KEY is not set in the environment variable. Use `echo $DEEPL_API_KEY` to check.
