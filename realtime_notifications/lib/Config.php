<?php

namespace OCA\RealTimeNotifications;


use OCP\IConfig;

class Config {

	const CFG_EVENT_SECRET = 'realtime_notifications.secret';
	const CFG_EVENT_BACKEND_URL = 'realtime_notifications.backend_host';

	private $config;

	public function __construct(IConfig $config) {
		$this->config = $config;
	}


	public function getBackendHost() {
		return $this->config->getSystemValue(
			self::CFG_EVENT_BACKEND_URL, 'http://localhost:8080'
		);
	}


	public function getSecret() {
		return $this->config->getSystemValue(
			self::CFG_EVENT_SECRET, ''
		);
	}

}