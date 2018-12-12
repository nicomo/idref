# idref

A Go client for the Web Services available for [IdRef](https://www.idref.fr) - a service run by ABES in France to expose authority records from [Calames](http://www.calames.abes.fr) (EAD finding aids), [Sudoc](http://www.sudoc.abes.fr) (Union catalog) and https://www.theses.fr (French PhDs).

The web services are documented (French only afaik) at http://documentation.abes.fr/aideidrefdeveloppeur/index.html

## Status 

This project is experimental/personal, use at your own risk.

## Covered

- get authority/authorities: can be used to retrieve a single auth record, or a bunch of auth records at a time. Currently covers Persons and Orgs only
- references: get the document associated with an authority in the IdRef database
- solr search: defaults to the "all" index if no known index is provided. So far only parses back Persons and Orgs records
- id2idref: give an ID from say Wikidata, retrieve an IdRef internal ID if it exists
- idref2id: give an IdRef internal ID, retrieve other know IDs. Does not use the "subservices" to limit the scope of the returned sources, _id est_ retrieves all the available IDs from all the available sources

See the documentation: https://godoc.org/github.com/nicomo/idref

## Not covered

- biblio: use references instead, which provides the same results (records from SUDOC) and more
- merged
- merged_inv
- iln2rcr
- iln2td3
