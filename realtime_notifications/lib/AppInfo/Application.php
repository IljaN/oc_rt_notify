<?php
namespace OCA\RealTimeNotifications\AppInfo;


use OCA\RealTimeNotifications\Authentication\TokenGenerator;
use OCA\RealTimeNotifications\Controller\AuthController;
use OCA\RealTimeNotifications\Controller\SettingsController;
use OCA\RealTimeNotifications\EventRouterService;
use OCA\RealTimeNotifications\ShareHooks;
use OCP\AppFramework\App;
use OCP\AppFramework\IAppContainer;
use GuzzleHttp\Client as GuzzleClient;

class Application extends App {

	const APP_ID = 'realtime_notifications';
	
	private $secret = '186163c9826c3a0762319a81a3889dd9';

	/**
	 *
	 * @param array $urlParams
	 */
	public function __construct(array $urlParams = []) {

		parent::__construct(self::APP_ID, $urlParams);


		$c = $this->getContainer();
		$s = $c->getServer();
		$eventDispatcher = $s->getEventDispatcher();

		$eventDispatcher->addListener(
			'OCA\Files::loadAdditionalScripts', function() {
				\OCP\Util::addScript(self::APP_ID, 'spop');
				\OCP\Util::addStyle(self::APP_ID, 'spop');
			}
		);

		$c->registerService('EventRouterService', function (IAppContainer $c) use ($s, $eventDispatcher) {
			$er = new EventRouterService(
				$s->getConfig(),
				new GuzzleClient(),
				$s->getContentSecurityPolicyManager(),
				$s->getGroupManager()

			);
			

			ShareHooks::initialize($er);

			return $er;
		});


		$c->query('EventRouterService');

		$c->registerService('SettingsController', function (IAppContainer $c) use ($s) {
			return new SettingsController(
				self::APP_ID,
				$c->query('Request'),
				$c->query('EventRouterService')
			);
		});

		$c->registerService('EventAuthTokenGenerator', function (IAppContainer $c) use ($s) {
			return new TokenGenerator($this->secret);
		});


		$c->registerService('AuthController', function (IAppContainer $c) use ($s) {
			return new AuthController(
				self::APP_ID,
				$c->query('Request'),
				$s->getUserSession(),
				$c->query('EventAuthTokenGenerator')
			);
		});
	}
}
