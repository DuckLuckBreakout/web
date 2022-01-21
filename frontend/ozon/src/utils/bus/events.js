const Events = {
    DummyEvent: 'dummy-event',

    LoginSendData: 'login-send-data',
    LoginEmitResult: 'login-emit-result',
    LoginIncorrectForm: 'login-incorrect-form',

    SignupSendData: 'signup-send-data',
    SignupEmitResult: 'signup-emit-result',
    SignupIncorrectForm: 'signup-incorrect-form',

    ProfileNewUserLoggedIn: 'profile-new-user-logged-in',
    ProfileFLNameChange: 'profile-flname-change',
    ProfileFLNameResult: 'profile-flname-result',
    ProfileIncorrectFLName: 'profile-incorrect-flname',
    ProfileAvatarChange: 'profile-avatar-change',
    ProfileAvatarResult: 'profile-avatar-result',
    ProfileIncorrectAvatar: 'profile-incorrect-avatar',
    ProfileEmailResult: 'profile-email-result',
    ProfileCheckAuthResult: 'profile-check-auth-result',
    ProfileAllResult: 'profile-all-result',
    ProfileAllGet: 'profile-all-get',
    ProfileLogoutPrepare: 'profile-logout-prepare',
    ProfileLogout: 'profile-logout',
    ProfileLogoutEmitResult: 'profile-logout-emit-result',
    ProfileIncorrectLogout: 'profile-logout-incorrect',
    ProfileTransmitData: 'profile-transmit-user-data',

    ProductsLoad: 'products-load',
    ProductsLoadSearch: 'products-load-search',
    ProductsLoaded: 'products-loaded',
    ProductsChangeCategory: 'products-change-category',
    ProductsItemAdded: 'products-item-added',
    RenderProductReviews: 'render-product-reviews',
    ProductTransmitData: 'product-transmit-data',
    ProductsItemNotAdded: 'products-item-removed',
    ProductStarsCounterLoad: 'product-stars-counter-load',
    ProductStarsCounterLoaded: 'product-stars-counter-loaded',
    ProductsCartLoadedProductsID: 'products-cart-loaded-products-ids',


    ProductChangeID: 'product-changeID',
    ProductLoad: 'product-load',
    ProductLoaded: 'product-loaded',
    ProductItemNotAdded: 'product-item-not-added',
    ProductItemAdded: 'product-item-added',
    ProductCartLoadedProductsID: 'product-cart-loaded-products-ids',
    ProductRenderReviewButton: 'product-render-review-button',

    HeaderLoad: 'product-load',
    HeaderLoaded: 'product-loaded',
    HeaderChangeCatalogID: 'header-change-catalog-id',
    HeaderChangeCategoryID: 'header-change-category-id',
    HeaderChangeCartItems: 'header-change-cart-items',
    HeaderSetCartItems: 'header-set-cart-items',

    CartAddProduct: 'cart-add-product',
    CartProductChange: 'cart-update-product',
    CartRemoveProduct: 'cart-remove-product',
    CartProductRemoved: 'cart-removed-product',
    CartLoad: 'cart-load',
    CartLoaded: 'cart-loaded',
    CartProductAdded: 'cart-product-added',
    CartLoadProductsAmount: 'cart-load-products-amount',
    CartLoadedProductsAmountReaction: 'cart-loaded-products-amount-reaction',
    CartGetProductsID: 'cart-get-products-ids',
    CartGetProductID: 'cart-get-product-id',
    CartLoadedProductID: 'cart-loaded-product-id',
    CartAddLastProduct: 'cart-add-last-product',
    CartDrop: 'cart-drop',
    CartCheckRecommendationsInCart: 'cart-check-recommendations-in-cart',

    OrderLoad: 'order-load',
    OrderLoaded: 'order-loaded',
    OrderNewBill: 'order-new-bill',
    OrderIncorrectPromo: 'order-incorrect-promo',
    SendOrder: 'order-send',

    OrdersLoad: 'orders-load',
    OrdersLoadMoreOrders: 'orders-load-more-orders',
    OrdersMoreOrdersLoaded: 'orders-more-orders-loaded',
    OrdersLoaded: 'orders-loaded',
    ReviewRightsLoad: 'review-rights-load',
    ReviewRightsLoaded: 'review-rights-loaded',
    ReviewUserDataLoad: 'review-user-data-load',
    ReviewUserDataLoaded: 'review-user-data-loaded',
    ReviewProductDataLoad: 'review-product-data-load',
    ReviewProductDataLoaded: 'review-product-data-load',
    ReviewOrder: 'review-send',
    GetProductReviews: 'get-product-reviews',
    GetProductReviewsReaction: 'get-product-reviews-reaction',

    ChangeReviewProductId: 'change-review-product-id',
    OrderSent: 'order-sent',

    SendPromo: 'send-promo',
    PromoSent: 'promo-sent',
    RecommendationLoad: 'recommendation-load',
    RecommendationLoaded: 'recommendation-loaded',

    WebPushSubscribe: 'web-push-subscribe',
    WebPushUnsubscribe: 'web-push-unsubscribe',
};


export default Events;