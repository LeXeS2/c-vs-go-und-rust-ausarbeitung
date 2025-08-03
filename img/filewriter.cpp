struct FileWriter {
    FILE* file;

    ~FileWriter() {
        if (file) fclose(file);
    }
};
