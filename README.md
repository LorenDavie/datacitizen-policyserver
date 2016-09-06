# datacitizen-policyserver

This is the code for the **Policy Server**, which governs the approval or disapproval of data transfers based on citizen data policies.

## Operations

### Data Transfer Request
A DTR is a request to transfer records to a Data Consumer. A DTR must define:

* The classification of the data being transferred (example: *e-commerce/books*).
* The identity of the Data Consumer to receive the data.
* A list of identities of the Data Citizens for whom the records pertain.

The Policy server will respond to the DTR with a *scrub list*, i.e. the list of records that must be scrubbed from the transfer.

#### DTR Request Format (JSON)

	{
		"producer":"Bob's Bookstore@@www.bobsbookstore.com",
		"consumer":"Consumer Research Trends@@www.crtrends.info",
		"data-category":"e-commerce/books",
		"citizen-identities":[
			"loren@axilent.com@@email",
			"bonzo@clownsforcrime.org@@email",
			"AxilentEngine@@Twitter"
		]
	}
	
#### DTR Response Format (JSON)

Some records must be scrubbed:

	{
		"status":"CONDITIONAL"
		"scrub-list":[
			"loren@axilent.com@@email",
			"bonzo@clownsforcrime.org@@email"
		]
	}

No scrubbing necessary, the entire transfer set is fine:

	{
		"status":"OK"
	}

The entire transfer is disallowed (usually indicating something wrong with the status of the consumer):

	{
		"status":"DISALLOWED"
	}

Or finally, indicating something is wrong wth the DTR itself:

	{
		"status":"ERROR",
		"message":"Unrecognized data category."
	}
