<?php

namespace OCA\RealTimeNotifications;


use OC\Share\Constants;
use OCP\AppFramework\Http\EmptyContentSecurityPolicy;
use OCP\Http\Client\IClient;
use OCP\IConfig;
use OCP\IGroupManager;
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

	/** @var IGroupManager  */
	private $groupManager;


	private $backendHost;


	/**
	 * EventRouterService constructor.
	 *
	 * @param IConfig $cfg
	 * @param GuzzleClient $http
	 * @param IContentSecurityPolicyManager $cspManager
	 */
	public function __construct(IConfig $cfg, GuzzleClient $http, IContentSecurityPolicyManager $cspManager, IGroupManager $gm) {
		$this->cfg = $cfg;
		$this->http = $http;
		$this->cspManager = $cspManager;
		$this->groupManager = $gm;

		$this->backendHost = $cfg->getSystemValue(
			self::CFG_EVENT_BACKEND_URL,
			'http://localhost:8080'
		);


		$this->cspManager
			->addDefaultPolicy(
				(new EmptyContentSecurityPolicy())
					->addAllowedConnectDomain($this->backendHost)
			);
	}

	public function onPostShare($params)  {
		$publishingEndpoint = $this->getPublishingEndpoint();
		$shareWith  = $params['shareWith'];
		$shareType = $params['shareType'];

		$event = ['to' => [], 'data' => []];

		switch ($shareType) {
			case Constants::SHARE_TYPE_USER:
				$event = ['to' => [$shareWith]];
				break;
			case Constants::SHARE_TYPE_GROUP:
				$shareWithGroup = $this->groupManager->get($shareWith);
				foreach ($shareWithGroup->getUsers() as $member) {
					$event['to'][] = $member->getUID();
				}
				break;
			default:
				return;
		}

		$event['data'] = $params;
		$this->http->post($publishingEndpoint, ['json' => $event]);
	}


	private function getPublishingEndpoint() {
		return $this->backendHost . '/events';
	}


	public function getBackendHost() {
		return $this->backendHost;
	}
}

