fn read_file() -> Result<()> {
    File::open("does_not_exist.txt")
        .context("Fehler beim Oeffnen der Datei")?;
    Ok(())
}
