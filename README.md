# deepldoc

deepldoc is a tool that automatically translates your project's documentation into another language. It is written in Go and uses DeepL, connected via API, as the translation engine.

## How to Use

First, create an account for the DeepL API. After creating an account, obtain an API key.
Next, set the API key as an environment variable. Execute the following command in the terminal:

```sh
export DEEPL_API_KEY=your_api_key_here
```

To use deepldoc, follow these instructions:

```sh
deepldoc your_directory target_language file_extension
```

Where:

- `your_directory` is the path to the directory containing the documents you want to translate.
- `target_language` is the language code of the target language (e.g., 'ja'). If omitted, 'ja' is used by default.
- `file_extension` is the extension of the file to be translated (e.g., 'md'). If omitted, 'md' is used by default.

## deepl

The additional translation command `deepl` is used as follows:

```sh
deepl text target_language
``` 

Where:

- `text` is the text to be translated.
-  `target_language` is the language code of the target language.

### Notes

**deepldoc** translates files with the specified file extensions and copies files with non-relevant extensions as they are, preserving the original directory structure.
**deepldoc** and **deepl** will help you with your multilingual projects. Please check them thoroughly and give us feedback if you have any problems.