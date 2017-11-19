<?php

return ['routes' => [
	[
		'verb' => 'GET',
		'url' => '/settings',
		'name' => 'settings#getBackendHost'
	],

	[
		'verb' => 'GET',
		'url' => '/admin_settings',
		'name' => 'settings#getAdminSettings'
	],

	[
		'verb' => 'GET',
		'url' => '/token',
		'name' => 'auth#getAuthToken'
	]
]];
