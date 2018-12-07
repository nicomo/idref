# idref

A Go client for the Web Services available for [IdRef](https://www.idref.fr) - a service run by ABES in France to expose authority records from [Calames](http://www.calames.abes.fr) (EAD finding aids), [Sudoc](http://www.sudoc.abes.fr) (Union catalog) and https://www.theses.fr (French PhDs).

The web services are documented (French only afaik) at http://documentation.abes.fr/aideidrefdeveloppeur/index.html

## Covered

- get authority: currently covers Persons and Orgs only
- references
- solr search: defaults to the "all" index if no known index is provided. So far only parses Persons (index persname_t) and Orgs (index corpname_t)
- id2idref
- idref2id: does not use the "subservices" to limit the scope of the returned sources

## Not covered

- biblio: use references instead, which provides the same results (records from SUDOC) and more
- merged
- merged_inv
- iln2rcr
- iln2td3
