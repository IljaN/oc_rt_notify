<?php

namespace OCA\RealTimeNotifications;


use GuzzleHttp\Promise;
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
	private $gm;


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
		$this->gm = $gm;

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

		if ($params['shareType'] == Constants::SHARE_TYPE_USER) {
			$this->http->post($this->backendHost . '/events', ['json' => $event]);
			return;
		}

		if ($params['shareType'] == Constants::SHARE_TYPE_GROUP) {
			$shareWithGroup = $this->gm->get($params['shareWith']);

			$promises = [];
			foreach ($shareWithGroup->getUsers() as $member) {
				$event['to'] = $member->getUID();
				$promises[$member->getUID()] = $this->http->postAsync(
					$this->backendHost . '/events', ['json' => $event]
				);
			}

			\GuzzleHttp\Promise\unwrap($promises);

		}
	}


	public function getBackendHost() {
		return $this->backendHost;
	}
}
