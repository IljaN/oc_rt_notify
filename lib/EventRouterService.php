<?php

namespace OCA\RealTimeNotifications;


use OCP\AppFramework\Http\EmptyContentSecurityPolicy;
use OCP\Http\Client\IClient;
use OCP\IConfig;
use OCP\Security\IContentSecurityPolicyManager;
use GuzzleHttp\Client as GuzzleClient;


class EventRouterService {

	const CFG_EVENT_BACKEND_URL = 'realtime_notifications.backend_host';

	/** @var  IConfig */
	private $cfg;

	/** @var  IClient */
	private $http;

	/** @var IContentSecurityPolicyManager  */
	private $cspManager;


	private $backendHost;


	/**
	 * EventRouterService constructor.
	 *
	 * @param IConfig $cfg
	 * @param GuzzleClient $http
	 * @param IContentSecurityPolicyManager $cspManager
	 */
	public function __construct(IConfig $cfg, GuzzleClient $http, IContentSecurityPolicyManager $cspManager) {
		$this->cfg = $cfg;
		$this->http = $http;
		$this->cspManager = $cspManager;

		$this->backendHost = $cfg->getSystemValue(
			self::CFG_EVENT_BACKEND_URL,
			'http://localhost:8080'
		);
	}

	public function run() {
		$this->cspManager->addDefaultPolicy((new EmptyContentSecurityPolicy())
			->addAllowedConnectDomain($this->backendHost)
		);
	}

	public function onPostShare($params)  {
		$event = ['to' => $params['shareWith'], 'data' => $params];
		$this->http->post($this->backendHost . '/events', ['json' => $event]);
	}


	public function getBackendHost() {
		return $this->backendHost;
	}
}
