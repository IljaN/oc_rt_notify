<?php

namespace OCA\RealTimeNotifications;


use OC\Share\Constants;
use OCA\RealTimeNotifications\Authentication\TokenGenerator;
use OCP\AppFramework\Http\EmptyContentSecurityPolicy;
use OCP\Http\Client\IClient;
use OCP\IConfig;
use OCP\IGroupManager;
use OCP\Security\IContentSecurityPolicyManager;
use GuzzleHttp\Client as GuzzleClient;


class EventRouterService {



	/** @var  Config */
	private $cfg;

	/** @var  IClient */
	private $http;

	/** @var IContentSecurityPolicyManager  */
	private $cspManager;

	/** @var IGroupManager  */
	private $groupManager;

	/** @var TokenGenerator  */
	private $tokenGenerator;


	private $backendHost;



	/**
	 * EventRouterService constructor.
	 *
	 * @param Config $cfg
	 * @param GuzzleClient $http
	 * @param IContentSecurityPolicyManager $cspManager
	 */
	public function __construct(Config $cfg, GuzzleClient $http, IGroupManager $gm, TokenGenerator $tokenGenerator) {
		$this->cfg = $cfg;
		$this->http = $http;
		$this->groupManager = $gm;
		$this->tokenGenerator = $tokenGenerator;
		$this->backendHost = $cfg->getBackendHost();
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

		$this->http->post($publishingEndpoint, [
			'headers' => ['Authorization' => "BEARER {$this->tokenGenerator->generatePublisherToken()}"],
			'json' => $event
		]);
	}


	private function getPublishingEndpoint() {
		return $this->backendHost . '/events';
	}


	public function getBackendHost() {
		return $this->backendHost;
	}
}

