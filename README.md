Genomic data explorer

---

High-level development log:
* built basic backend (golang + postgres database deployed via docker)
    * restructured backend subdirectory to be more consistent with golang best practices
* realized `Python` makes a lot more sense for the backend:
    * `Python` has a lot of libraries for working with genomic/biological data
    * `Python` makes it easier to interact with both bioinformatics and LLM ecosystems
    * consistency with ETL pipeline would be a plus
* archived golang backend version of the project
