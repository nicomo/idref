# idref

## DO NOT USE - Work in progress

A Go client for the Web Services available for [IdRef](https://www.idref.fr) - a service run by ABES in France to expose authority records from [Calames](http://www.calames.abes.fr) (EAD finding aids), [Sudoc](http://www.sudoc.abes.fr) (Union catalog) and https://www.theses.fr (French PhDs).

The web services are documented (French only afaik) at http://documentation.abes.fr/aideidrefdeveloppeur/index.html

## Covered

- get authority: retrieves an authority record when you provide an ID for it (PPN). Currently covers Persons and Corporations only
- references: get the documents linked to an authority record when you provide an ID for it (PPN)
- solr search
  - search a Person (index persname_t)
  - search an Organization (index corpname_t)

## Not covered

- biblio: use references instead
- merged
- merged_inv
- idref2id
- id2idref
- iln2rcr
- iln2td3
