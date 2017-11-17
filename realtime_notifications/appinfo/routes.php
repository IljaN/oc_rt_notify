<?php

return ['routes' => [
	[
		'verb' => 'GET',
		'url' => '/settings',
		'name' => 'settings#get'
	],

	[
		'verb' => 'GET',
		'url' => '/token',
		'name' => 'auth#getAuthToken'

	]
]];
