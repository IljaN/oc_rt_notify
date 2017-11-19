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

		$container = $this->getContainer();
		$server = $container->getServer();

		$server->getEventDispatcher()->addListener(
			'OCA\Files::loadAdditionalScripts', function() {
				\OCP\Util::addScript(self::APP_ID, 'spop');
				\OCP\Util::addStyle(self::APP_ID, 'spop');
			}
		);

		$container->registerService('EventAuthTokenGenerator', function (IAppContainer $c) use ($server) {
			return new TokenGenerator($this->secret);
		});

		$container->registerService('EventRouterService', function (IAppContainer $c) use ($server) {
			$er = new EventRouterService(
				$server->getConfig(),
				new GuzzleClient(),
				$server->getContentSecurityPolicyManager(),
				$server->getGroupManager(),
				$c->query('EventAuthTokenGenerator')
			);
			

			ShareHooks::initialize($er);

			return $er;
		});


		$container->query('EventRouterService');

		$container->registerService('SettingsController', function (IAppContainer $c) use ($server) {
			return new SettingsController(
				self::APP_ID,
				$c->query('Request'),
				$c->query('EventRouterService')
			);
		});

		$container->registerService('AuthController', function (IAppContainer $c) use ($server) {
			return new AuthController(
				self::APP_ID,
				$c->query('Request'),
				$server->getUserSession(),
				$c->query('EventAuthTokenGenerator')
			);
		});
	}
}
