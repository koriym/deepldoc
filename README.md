# A Command line tool for deepl


deepldoc is a tool that automatically translates your project's documentation into another language, written in the Go language and using DeepL connected via API as the translation engine.

## How to use

First, create an account for the DeepL API. After creating an account, obtain an API key.
Next, set the API key as an environment variable. Execute the following command in the terminal:

```sh
export DEEPL_API_KEY=your_api_key_here
```

## deepl

The additional translation command deepl is used as follows:

```
. /deepl text target_language
``` 

- **text** is the text to be translated.
- **target_language** is the language code of the target language.

## deepldoc

To use deepldoc, follow these instructions:

```sh
. /deepldoc your_directory target_language file_extension
``````

- **your_directory** is the path to the directory containing the documents you want to translate.
- **target_language** is the language code of the target language (e.g. 'ja'). If omitted, 'ja' is used by default.
- **file_extension** is the extension of the file to be translated (e.g. 'md'). If omitted, 'md' is used by default.


### Notes

It is important to specify the correct file path and language. If each is not specified correctly, an error will occur.
deepldoc translates files with the relevant file extensions and copies files with non-relevant extensions as they are. This preserves the original directory structure.
deepldoc and deepl will help you with your multilingual projects. Please check it thoroughly and give us feedback if you have any problems.

## How to Download and Run Release Assets

Follow these steps to download and run the `deepl` and `deepldoc` files:

1. Go to the [releases page](https://github.com/your-repository/releases) and download `deepl` and `deepldoc`.

2. Open a terminal and navigate to your downloads folder:
    ```sh
    cd ~/Downloads
    ```

3. Make the files executable:
    ```sh
    chmod +x deepl
    chmod +x deepldoc
    ```

4. Run the files:
    ```sh
    ./deepl
    ./deepldoc
    ```

That's it!
